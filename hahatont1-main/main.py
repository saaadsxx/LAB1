import requests
import streamlit as st
import pandas as pd
import os

# Настройки для API Twitch
TWITCH_CLIENT_ID = os.getenv('TWITCH_CLIENT_ID', 'wbcobfa0nhlbk2ay8catecsipd5jkt')
TWITCH_SECRET = os.getenv('TWITCH_SECRET', '3qwh9edliut8uwijwit8kw7l1ill9e')

# Функция для получения токена доступа Twitch
@st.cache_data(show_spinner=False)
def get_twitch_token():
    url = 'https://id.twitch.tv/oauth2/token'
    params = {
        'client_id': TWITCH_CLIENT_ID,
        'client_secret': TWITCH_SECRET,
        'grant_type': 'client_credentials'
    }
    response = requests.post(url, params=params)

    if response.status_code != 200:
        st.error(f"Ошибка при получении токена Twitch: {response.status_code}")
        return None

    data = response.json()
    return data.get('access_token')

# Функция для получения стримов с фильтрацией
@st.cache_data(show_spinner=False)
def get_twitch_streamers(limit=20, game=None, language=None):
    token = get_twitch_token()
    if token is None:
        return pd.DataFrame()

    headers = {
        'Authorization': f'Bearer {token}',
        'Client-Id': TWITCH_CLIENT_ID
    }
    url = 'https://api.twitch.tv/helix/streams'
    
    params = {'first': limit}
    if game:
        params['game_id'] = game
    if language:
        params['language'] = language

    response = requests.get(url, headers=headers, params=params)

    if response.status_code != 200:
        st.error(f"Ошибка при получении данных Twitch: {response.status_code}")
        return pd.DataFrame()

    data = response.json()
    if 'data' not in data:
        st.warning('Нет доступных данных о стримах.')
        return pd.DataFrame()

    streams = []
    for stream in data['data']:
        streams.append({
            'Streamer': stream['user_name'],
            'Game': stream['game_name'],
            'Viewers': stream['viewer_count'],
            'Language': stream['language'],
            'Started At': stream['started_at']
        })

    df = pd.DataFrame(streams)
    df.index = df.index + 1
    return df

# Функция для получения списка категорий игр (для фильтрации)
@st.cache_data(show_spinner=False)
def get_twitch_games():
    token = get_twitch_token()
    if token is None:
        return []

    headers = {
        'Authorization': f'Bearer {token}',
        'Client-Id': TWITCH_CLIENT_ID
    }
    url = 'https://api.twitch.tv/helix/games/top'
    response = requests.get(url, headers=headers)

    if response.status_code != 200:
        st.error(f"Ошибка при получении категорий игр Twitch: {response.status_code}")
        return []

    data = response.json()
    games = [{'name': game['name'], 'id': game['id']} for game in data['data']]
    return games

# Streamlit интерфейс
st.title('Анализ популярных стримеров на Twitch')

# Получение списка игр для фильтрации
games = get_twitch_games()
game_options = ['Все'] + [game['name'] for game in games]
selected_game = st.selectbox('Выберите игру для фильтрации:', game_options)

# Получение списка языков для фильтрации
language_options = ['Все', 'en', 'ru', 'fr', 'es']  # Добавьте другие языки по необходимости
selected_language = st.selectbox('Выберите язык стримов:', language_options)

# Количество стримеров для отображения
limit = st.slider('Количество стримов для отображения:', 5, 100, 20)

# Определение выбранной игры и языка для запроса
game_id = None if selected_game == 'Все' else next(game['id'] for game in games if game['name'] == selected_game)
language = None if selected_language == 'Все' else selected_language

# Twitch Streamers
st.header('Популярные стримеры на Twitch')
twitch_streamers_df = get_twitch_streamers(limit=limit, game=game_id, language=language)
st.dataframe(twitch_streamers_df)

# Кнопка для скачивания CSV
csv_data = twitch_streamers_df.to_csv(index=False).encode('utf-8')
st.download_button(
    label="Скачать CSV",
    data=csv_data,
    file_name='twitch_streamers.csv',
    mime='text/csv',
)

# Кнопка для скачивания JSON
json_data = twitch_streamers_df.to_json(orient='records').encode('utf-8')
st.download_button(
    label="Скачать JSON",
    data=json_data,
    file_name='twitch_streamers.json',
    mime='application/json',
)

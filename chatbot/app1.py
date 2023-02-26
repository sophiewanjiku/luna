# Deploy the app to Streamlit Sharing

import streamlit as st
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns

st.title('Mental Health Chatbot')

st.write('This is a chatbot that will help you to understand your mental health and provide you with resources to help you.')

st.write('Feel free to ask any questions you may have about mental health.')

st.write('Please note that this chatbot is not a substitute for professional help.')

# Load the model and tokenizer
import pickle
import tensorflow as tf
from tensorflow.keras.preprocessing.text import Tokenizer
from tensorflow.keras.preprocessing.sequence import pad_sequences
from tensorflow.keras.models import load_model

model = pickle.load(open('/home/ashioyajotham/luna/model.pkl', 'rb'))
tokenizer = pickle.load(open('/home/ashioyajotham/luna/tokenizer.pkl', 'rb'))

# The aim is to create a chatbot that will help the user to understand their mental health and provide them with resources to help them.
# The chatbot will be able to answer questions about mental health and provide resources to help the user.

st.write('Please enter your question below:')
question = st.text_input('')
question = question.lower()

# Create a function to predict the answer
def predict_answer(question):
    seq = tokenizer.texts_to_sequences([question])
    padded = pad_sequences(seq, maxlen=100)
    pred = model.predict(padded)
    pred = np.argmax(pred)
    return pred

# Create a function to get the answer
def get_answer(pred):
    # Generate random answer
    return random.choice(answers[pred])

st.write('Answer:')
st.write(get_answer(predict_answer(question)))
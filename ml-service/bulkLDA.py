"""
This script is used to perform LDA on the question and answer columns of a CSV file 
from ../data/record/. 
The CSV file is read into a DataFrame, and then each row is processed. 
The question and answer columns are tokenized and stop words are removed. 
The resulting lists of tokens are then used to perform LDA. 
The results are then written to a new CSV file.

noted if the filed is empty it will return empty list
"""
import os
import pandas as pd
from app.services.tokenize_service import TokenizeService
from app.services.stopWords_service import StopWordsService
from app.services.lda_service import LDAServerice

def replace_chars(text: str) -> str:
    """Replace unwanted characters in a string."""
    if pd.isnull(text):
        return ""
    text = text.replace('"', '')
    text = text.replace('\n', '')
    return text

def tokenize_and_remove_stop_words(text: str) -> list:
    """Tokenize a string and remove stop words."""
    tokens = TokenizeService.tokenize_from_string(text)
    bagOfWords = StopWordsService.remove_stop_words_from_list(tokens)
    return bagOfWords

def lda_text(text: str,num_topics :int) -> list:
    if text == "":
        return []
    replaceCharText = replace_chars(text)
    tokenizeText = tokenize_and_remove_stop_words(replaceCharText)
    ldaResult = LDAServerice.perform_lda(tokenizeText,num_topics)
    return ldaResult

## -------

csv_files = [f for f in os.listdir('../data/record/') if f.endswith('.csv')]

for csv_file in csv_files:
    df = pd.read_csv('../data/record/' + csv_file)
    
    new_rows = []

    num_topics = 5
    zero_indices = []

    # Loop over each row in the DataFrame
    for _, row in df.iterrows():
        index = row['youtubeURL'] + "-" + str(row['index'])
        print(index, "Start....  Result:",end="")
        
        question_lda = lda_text(row['question'],num_topics)
        answer_lda = lda_text(row['answer'],num_topics)

        if question_lda == [0]  or answer_lda == [0] :
            zero_indices.append(index)

        # Create a new row with the index, question, and answer
        new_rows.append({'index': index, 'question': question_lda, 'answer': answer_lda})
        print("Success")

    # Create a new DataFrame with the new rows
    new_df = pd.DataFrame(new_rows)

    # Write the new DataFrame to a CSV file
    # Write the new DataFrame to a CSV file
    new_filename = csv_file.replace('.csv', '-lda.csv')
    new_df.to_csv(f'../data/lda/{new_filename}', index=False)
    print("Indices that returned [0] * num_topics:", zero_indices)
    print(f'Wrote {len(new_df)} rows to ../data/lda/{new_filename}')


## Contribute data
# 1. Download data from /data/record [raw data]
# 2. Process csv (vec) data in format of [index, question, answer]
# 3. Upload to /data/text2vec with name of <model>.csv
# 4. prepare apis for each model
# 5. Update config file for each model
# 6. Run workflow to deploy updated
# 7. Refresh service in server use makefile
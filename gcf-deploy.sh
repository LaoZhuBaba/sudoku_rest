gcloud functions deploy $1 \
    --gen2 \
    --runtime=go121 \
    --region=$2 \
    --source=sudokurest \
    --entry-point SudokuRest \
    --trigger-http \
    --allow-unauthenticated 

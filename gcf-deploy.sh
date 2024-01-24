  gcloud functions deploy sudokurest \
    --gen2 \
    --runtime=go121 \
    --region=australia-southeast1 \
    --source=sudokurest \
    --entry-point SudokuRest \
    --trigger-http \
    --allow-unauthenticated 

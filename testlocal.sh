url=http://localhost:8080/SudokuRest
echo Testing locally using URL: $url
curl -i -X POST -d @$1 $url
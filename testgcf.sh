url=$(gcloud functions describe --region $2 $1 --format="get(url)")
echo Testing GCF using URL: $url
curl -i -X POST -d @$3 $url

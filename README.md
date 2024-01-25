Deploy Google Cloud function which solves Sudoku puzzles via a REST interface.

This code supports 4x4, 9x9 (normal) and 16x16 puzzles.  16x16 works but takes several minutes to complete when run as a GCF so you probably need to increase the HTTP timeout and maybe increase memory.

Example JSON input can be found in: sample2.json (4x4), sample3.json (9x9) and sample4.json (16x16).  Successful output is in the same format but with puzzle completed.

Note that the Sudoku data is given as a one dimensional array.  Visualising this as a grid makes solving easier for humans but the code logic works one-dimensionally.

The RelatedElements table is generated automatically by a Python script.  Three different tables are produced (one for each supported puzzle size) and then combined.  Use "make lookup_table" to generate the lookup_table.go file. 

Three 9x9 examples in a raw format are in samples.txt.  These are from https://projecteuler.net/problem=96

The data in the sample4.json was taken from here: https://gist.github.com/vaskoz/8212615.

The code is in the sudokurest directory with a Makefile which provides targets for deploying to GCP, running locally, running tests, etc.

"make deploy" just runs gcf-deploy.sh.  This deploys a Cloud Function to your currently configured GCP project.  Or you can specify by adding --project=project_name 

Test the REST interface either locally or the GCF with: "make test_local" and "make test_gcf".  For local testing you need to start the local server first.

Todo: write tests for 4x4 & 16x16 and test failure modes.
Deploy Google Cloud function which solves 9x9 Sudoku puzzles via a REST interface.

Example JSON input can be found in: sample.json.  Successful output is in the same format but with puzzle completed.

Note that the Sudoku data is given as a one dimensional array.  Visualising this as a grid makes solving easier for humans but the code logic works one-dimensionally.

The RelatedElements table is generated automatically by a Python script.  I don't intend to generalise the solver to work with other Shapes (e.g., 16x16) but I think the solution should generalise reasonably easily if you built a RelatedElements table on the fly for each Sudoku shape.  This is why I have included the "length" and "maxValue" parameters in the JSON format.  E.g., for a 16x16 puzzle you would need a different RelatedElements table and "length" equal to 256 and "maxValue" of "F" (assuming you standard hexadecimal digits)

Three examples in a raw format are in samples.txt.  These are from https://projecteuler.net/problem=96

The code is in the sudokurest directory with a Makefile which provides targets for deploying to GCP, running locally, running tests, etc.

"make deploy" just runs gcf-deploy.sh.  This deploys a Cloud Function to your currently configured GCP project.  Or you can specify by adding --project=project_name 

Test the REST interface either locally or the GCF with: "make test_local" and "make test_gcf".  For local testing you need to start the local server first.

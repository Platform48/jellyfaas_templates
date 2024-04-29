# f_pound_to_dollar
Sample FaaS function to convert a pound to a dollar.

``zip -x "*.git*" "*.idea*" -r f_go_pound_to_dollar.zip f_go_pound_to_dollar``

Then you can upload or deploy from the command line

``gcloud functions deploy ConvertGBPToUSD --runtime go121 --trigger-http --allow-unauthenticated --region europe-west2 --gen2``
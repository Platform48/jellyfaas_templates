# f_php_pound_to_dollar
php version of £ to $

zip -x "*.git*" -r f_php_pound_to_dollar.zip f_php_pound_to_dollar

composer require google/cloud-functions-framework

export FUNCTION_TARGET=convertGBPToUSD

php -S localhost:8080 vendor/google/cloud-functions-framework/router.php
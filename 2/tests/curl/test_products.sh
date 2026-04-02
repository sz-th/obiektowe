#!/bin/bash

echo "--- GET ALL PRODUCTS ---"
curl -X GET http://localhost:8000/api/product
echo -e "\n"

echo "--- CREATE NEW PRODUCT (POST) ---"
curl -X POST http://localhost:8000/api/product \
     -H "Content-Type: application/json" \
     -d '{"name": "Testowy Produkt", "price": 19.99, "description": "Opis produktu"}'
echo -e "\n"
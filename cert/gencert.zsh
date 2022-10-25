#!/bin/zsh

CERT_DIR=./cert

rm ${CERT_DIR}/*.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ${CERT_DIR}/ca-key.pem -out ${CERT_DIR}/ca-cert.pem -subj "/C=FR/ST=Occitanie/L=Toulouse/O=Tech School/OU=Education/CN=*.techschool.guru/emailAddress=techschool.guru@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in ${CERT_DIR}/ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout ${CERT_DIR}/server-key.pem -out ${CERT_DIR}/server-req.pem -subj "/C=FR/ST=Ile de France/L=Paris/O=PC Book/OU=Computer/CN=*.pcbook.com/emailAddress=pcbook@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in ${CERT_DIR}/server-req.pem -days 60 -CA ${CERT_DIR}/ca-cert.pem -CAkey ${CERT_DIR}/ca-key.pem -CAcreateserial -out ${CERT_DIR}/server-cert.pem -extfile ${CERT_DIR}/server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in ${CERT_DIR}/server-cert.pem -noout -text

# Verify the certificate
openssl verify -CAfile ${CERT_DIR}/ca-cert.pem ${CERT_DIR}/server-cert.pem

# API REST para Upload de Arquivos no Amazon S3

Esta é uma API REST simples para facilitar o upload de arquivos no Amazon S3. A API permite que os usuários enviem arquivos para um bucket do Amazon S3 com facilidade, fornecendo um caminho de arquivo personalizado.


**Tecnologias Utilizadas**
* Golang: A linguagem de programação Go foi escolhida para desenvolver a API devido à sua eficiência, desempenho e facilidade de uso.
* Gin: O framework web Gin foi usado para simplificar o roteamento e o tratamento de solicitações HTTP.
* Amazon S3: O serviço de armazenamento de objetos da Amazon foi usado para armazenar os arquivos enviados pelos usuários.
* AWS SDK para Go: A biblioteca oficial da Amazon para Go foi usada para interagir com o Amazon S3.
* Viper: A biblioteca Viper foi utilizada para carregar configurações a partir de um arquivo .env, facilitando a configuração da aplicação.
* Zap: A biblioteca Zap foi usada para configuração de log robusta e eficiente.

**Funcionalidades Principais**

* Upload de arquivos para o Amazon S3.
* Personalização do caminho do arquivo no bucket S3.
* Retentativas em caso de falha no upload do arquivo.

**Configuração**

Certifique-se de configurar corretamente as variáveis de ambiente no arquivo .env no diretório /cmd/uploader. Você pode usar o seguinte formato como exemplo:

* AWS_REGION=us-east-1
* AWS_ACCESS_KEY=seu-access-key
* AWS_SECRET_KEY=seu-secret-key
* AWS_S3_BUCKET=seu-bucket-s3
* LOG_OUTPUT=stdout
* LOG_LEVEL=info
* GIN_MODE=release

**Como Executar**

Clone este repositório:

* Instale as dependências:
```
go mod tidy
```

* No diretório /cmd/uploader execute o seguinte comando para iniciar a aplicação:
```
go run main.go
```

**A API estará disponível em http://localhost:8080**

**Uso**

Envie uma solicitação POST para http://localhost:8080/api/v1/upload com o arquivo que deseja enviar. Certifique-se de incluir o campo "path" no corpo da solicitação para personalizar o caminho do arquivo no bucket S3.

**Exemplo de corpo da solicitação:**
```javascript
var formdata = new FormData();
formdata.append("path", "user/avatar");
formdata.append("files", fileInput.files[0], "download.png");

var requestOptions = {
  method: 'POST',
  body: formdata,
  redirect: 'follow'
};

fetch("http://localhost:8080/api/v1/upload", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

**Contribuição**

Sinta-se à vontade para contribuir com melhorias para este projeto. Se você encontrar problemas ou tiver sugestões, abra uma issue.
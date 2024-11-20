
System of stream videos

To separate the user in two groups (stream - manager).

for the Manager user the system must allow.
1. upload video.
2. delete video.
3. register new user of the type manager

for the stream group the system must allow.
1. interface web and mobile
2. Suggest videos based on categories videos and similarity user watch 
3. Suggest based on the taste of ordinary users are watching
4. watch videos
5. rate the video positively, regularly and negatively




Recomende vídeos com base em interesses do usuário (IA).
Permita ao usuario realizar upload de videos e gravar na API.
Permita que um 

Controle o acesso de vídeos, garantindo que cada usuário não possa assistir ao mesmo vídeo em duas sessões simultâneas.

Armazene as informações de cada vídeo e seu estado de visualização no banco de dados, com informações sobre progresso e data de última visualização.
Especificações do Desafio:

1. Inteligência Artificial para Recomendação
Implemente um sistema de recomendação que, ao final de cada vídeo, sugira o próximo vídeo baseado em:
Histórico de vídeos assistidos pelo usuário.
Preferências de outros usuários com comportamentos similares.
Classificação de vídeos feita pelo usuário (gostos/desgostos).
Utilize um modelo de filtragem colaborativa (Collaborative Filtering) ou um modelo baseado em conteúdo (Content-Based Filtering) para as recomendações.
2. Controle de Sessões de Acesso
No backend, implemente uma funcionalidade para impedir que um mesmo usuário possa abrir o vídeo em mais de uma sessão ao mesmo tempo.
Quando um usuário inicia a visualização de um vídeo, armazene essa sessão no banco de dados e bloqueie novas tentativas de acesso ao mesmo vídeo enquanto essa sessão estiver ativa.
Se o usuário tentar abrir o vídeo em outra aba ou dispositivo, o sistema deve retornar uma mensagem indicando que o vídeo já está em reprodução em outra sessão.
3. Armazenamento e Controle do Estado de Visualização
No banco de dados, armazene o progresso de cada vídeo, com informações de:
Data de início e último ponto assistido.
Total assistido (em segundos ou frames).
Status (completo ou incompleto).
Essas informações devem ser atualizadas em tempo real ou a cada período (e.g., a cada 10 segundos) enquanto o usuário assiste ao vídeo, para que o sistema retome a reprodução do ponto onde ele parou em uma sessão futura.
Requisitos Técnicos
Tecnologia Backend: Use Node.js, Python (Flask ou Django), ou outra linguagem/framework com suporte a manipulação de vídeo e inteligência artificial.
Banco de Dados: SQL ou NoSQL. Armazene os metadados dos vídeos e sessões no banco de dados.
IA e Recomendação: Utilize uma biblioteca de machine learning, como Scikit-Learn, TensorFlow, ou PyTorch, para implementar o sistema de recomendação.
Biblioteca de Controle de Sessão: No backend, use WebSocket ou uma implementação similar para acompanhar o progresso dos vídeos em tempo real.
Tarefas e Objetivos:
Implementar Recomendação de Vídeo:

Treinar um modelo de recomendação simples que sugira vídeos com base no histórico e nas preferências do usuário.
Controle de Sessão de Vídeo:

Desenvolver o controle de sessões para bloquear o acesso simultâneo do mesmo vídeo por um único usuário.
Salvar Progresso do Vídeo:

Implementar a gravação de progresso no banco de dados e a recuperação para continuidade em futuras sessões.
Testar e Validar Funcionalidades:

Certifique-se de que as funcionalidades de recomendação e controle de sessão estejam operando corretamente.
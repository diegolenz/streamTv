
"use client"; 

import React, { useState } from 'react';
import axios from 'axios';

 const VideoUpload = () => {
  const [selectedFile, setSelectedFile] = useState(null);

  const handleFileChange = (event) => {
    setSelectedFile(event.target.files[0]);
  };

  const handleUpload = async () => {
    if (!selectedFile) {
      alert('Por favor, selecione um arquivo para fazer o upload');
      return;
    }

    const formData = new FormData();
    formData.append('file', selectedFile);

    try {
      const response = await axios.post('http://localhost:8000/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      alert('Upload bem-sucedido!');
      console.log('Nome do arquivo salvo:', response.data.filename);
    } catch (error) {
      console.error('Erro no upload:', error);
      alert('Falha ao fazer o upload do vídeo');
    }
  };

  return (
    <div>
      <h2>Upload de Vídeo</h2>
      <input type="file" accept="video/mp4" onChange={handleFileChange} />
      <button onClick={handleUpload}>Fazer Upload</button>
    </div>
  );
};

export default VideoUpload;
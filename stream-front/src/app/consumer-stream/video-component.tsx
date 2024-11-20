"use client"; 
import React, { useEffect, useState } from 'react';


const VideoPlayer = ({ videoBytes }) => {
  const [videoUrl, setVideoUrl] = useState(null);

  useEffect(() => {
      if (videoBytes) {
          
        
          const blob = new Blob([videoBytes], { type: 'video/mp4' }); 
          const url = URL.createObjectURL(blob);
          setVideoUrl(url);

          
          return () => URL.revokeObjectURL(url);
      }
  }, [videoBytes]);

  if (!videoUrl) return <p>Carregando o vídeo...</p>;

  return (
      <video controls width="600" src={videoUrl} />
  );
};

export default VideoPlayer;
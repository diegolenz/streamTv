"use client";
import axios from "axios";
import { useEffect, useState } from "react";
import VideoPlayer from "./video-component";

interface VideoChuck {
    id?: string;
    fileName?: string;
    Chunk?: any;// ArrayBuffer;
    index?: number;
}

export default function ConsumerVideo() {

    // Array para armazenar os chunks recebidos

    const [videos, setVideos] = useState(null);

    const consumerVideo = async () => {
        const chunks = [];
        const totalChunck = 2;
        let bufferTotal: Uint8Array;
        try {
            for (let chunkIndex = 0; chunkIndex < totalChunck; chunkIndex++) {

                const response = await axios.get(`http://localhost:8000/find-video/${chunkIndex}`);

                if (response.status != 200) {
                    throw new Error(`Erro na chamada da API: ${response.statusText}`);
                }
                const data: VideoChuck = response.data;
                const binaryString = window.atob(data.Chunk);
                const len = binaryString.length;

                let ini = 0;
                let lenTotal = len;
                if (bufferTotal != null) {
                    ini = bufferTotal.length;
                    lenTotal += bufferTotal.length;
                }
                console.log(lenTotal)
                let bufferAtual = new Uint8Array(lenTotal);

                if (bufferTotal != null) {
                    console.log('carregando buffer antigo')
                    for (let i = 0; i < bufferTotal.length; i++) {
                        bufferAtual[i] = bufferTotal[i];
                    }
                }
               
              
                console.log('ini' + ini)
                let countAux = 0;
                for (let i = (ini); i < lenTotal; i++) {
                    bufferAtual[i] = binaryString.charCodeAt(countAux);
                    countAux += 1;
                }

                bufferTotal = bufferAtual;
                console.log(bufferTotal)

                //bufferTotal = new Uint8Array([...bufferTotal, ...buffer]);


                const byteArray = new Uint8Array(bufferAtual)
                chunks.push(byteArray);

                const combinedBlob = new Blob([byteArray], { type: 'video/mp4' });


                const url = URL.createObjectURL(combinedBlob);

                setVideos(url);

                /* videoElement.onplaying = () => {
                    fetchAndPlayChunk(index + 1); // Busca o próximo chunk
                }; */
            }

        } catch (err) {
            console.log(err)
            throw 'falha ao Buscar Video'

        }

    };

    useEffect(() => {
        consumerVideo();
    }, []);

    return (
        <div>
            <h1>Visualização de Vídeo</h1>
            <div>
                <h1>Exibição de Vídeo</h1>

                {videos ? <video controls width="600" src={videos} /> : <p>Carregando...</p>}
            </div>
        </div>
    );
}
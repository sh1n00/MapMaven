using System;
using UnityEngine;
using UnityEngine.Serialization;

namespace Types
{
    [Serializable]
    public class AudioResponse
    {
        [FormerlySerializedAs("audio_binary")]
        public string audio_binary;

        // public byte[] GetAudioBytes()
        // {
        //     return Convert.FromBase64String(audioBinary);
        // }
    }
}
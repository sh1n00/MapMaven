using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;
using System.IO;
using System.Text;
using static Wav;

[RequireComponent(typeof(AudioSource))]
public class AudioClipGenerate : MonoBehaviour
{
    AudioClip tmp;
    private AudioClip m_AudioClip;
    private float m_AudioLevel;
    [SerializeField] private string m_DeviceName;

    // サンプリング周波数
    private const int samplingFrequency = 44100;
    // 録音する秒数
    public int recordingTime = 5;

    public AudioSource audioSource;
    private AudioClip clip;

    private string targetDevice = "";

    private string outputFilePath = "Assets/output.wav";
    
    // Start is called before the first frame update
    void Start()
    {
        targetDevice = "";
        
        foreach (var device in Microphone.devices) {
            Debug.Log($"Device Name: {device}");
            if (device.Contains(m_DeviceName)) {
                targetDevice = device;
            }
        }
        
        Debug.Log($"=== Device Set: {targetDevice} ===");
        //m_AudioClip = Microphone.Start(targetDevice, true, 10, 48000);
    }

    // Update is called once per frame
    void Update()
    {
        if (Input.GetKeyDown(KeyCode.Space)) {
            Debug.Log("Down Space Key");
            // オーディオクリップを生成する
            clip = Microphone.Start(targetDevice, true, recordingTime, samplingFrequency);
            
            // 録音を開始する
            audioSource.clip = clip;
            audioSource.Play();
        }
        
        // 録音が終了した場合
        if (Input.GetKeyUp(KeyCode.Space)) {
            // 録音を停止する
            Microphone.End(null);
            Wav.ExportWav(clip, outputFilePath);
            Debug.Log("Export WAV");
        }
        
        
    }
    //
    

    
}

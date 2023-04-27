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
    private AudioClip m_AudioClip;
    private float m_AudioLevel;
    [SerializeField] private string m_DeviceName;

    // サンプリング周波数
    private const int samplingFrequency = 44100;
    

    public AudioSource audioSource;
    private AudioClip clip;

    private string targetDevice = "";
    private string outputFilePath = "Assets/output.wav";
    public static string wavBase64;

    public float sensitivity = 0.1f; // 録音を開始する音量の閾値
    public float maxRecordTime = 10f; // 最大録音時間
    public float minSilenceTime = 1f; // 録音を終了するための最小の無音時間
    
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
            clip = Microphone.Start(targetDevice, true, (int)maxRecordTime, samplingFrequency);
            
            // 録音を開始する
            audioSource.clip = clip;
            audioSource.Play();
        }
        
        // 録音が終了した場合
        if (Input.GetKeyUp(KeyCode.Space)) {
            // 録音を停止する
            Microphone.End(null);
            var wavByte =Wav.ToWav(clip);
            Wav.ExportWav(clip, outputFilePath);
            wavBase64 = Convert.ToBase64String(wavByte);
            Debug.Log(wavBase64);
            //AudioTranscription.AudioTrans(outputFilePath);
            Debug.Log("Export WAV");
        }
        
        
    }
    //
    

    
}

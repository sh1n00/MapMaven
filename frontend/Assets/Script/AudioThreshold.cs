using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Audio;
using System;
using System.IO;
using System.Text;
using static Wav;

public class AudioThreshold : MonoBehaviour
{
    public float sensitivity = 0.1f; // 録音を開始する音量の閾値
    public float maxRecordTime = 10f; // 最大録音時間
    public float minSilenceTime = 1f; // 録音を終了するための最小の無音時間
    public string outputFilePath = "Assets/output.wav";
    [SerializeField] private string m_DeviceName;

    private string targetDevice = "";
    private bool isRecording = false;
    private float[] samples; // マイクから取得したサンプルデータ
    private float silenceTime = 0f;
    private AudioClip recordedClip;
    // Start is called before the first frame update
    void Start()
    {
        //samples = new float[1024];
        // マイク入力の設定
        AudioSettings.GetDSPBufferSize(out var bufferLength, out var numBuffers);

        targetDevice = "";
        
        foreach (var device in Microphone.devices) {
            Debug.Log($"Device Name: {device}");
            if (device.Contains(m_DeviceName)) {
                targetDevice = device;
            }
        }
        
        recordedClip = Microphone.Start(null, true, (int)maxRecordTime, AudioSettings.outputSampleRate);
        //Microphone.End(null);
    }

    // Update is called once per frame
    void Update()
    {
        //float[] samples = new float[audioSource.clip.samples];; // マイクから取得したサンプルデータ
        samples = new float[1024];
        recordedClip.GetData(samples, 0);

        float sum = 0;
        float amplitude = 0;
        for (int i = 0; i < samples.Length; i++)
        {
            sum += Mathf.Abs(samples[i]);
        }
        //amplitude /= samples.Length;
        amplitude = Mathf.Clamp01(sum * 1f / (float)samples.Length);


        Debug.Log(amplitude);
        

        // 録音中に音量が閾値以下になった場合、無音時間を加算する
        if (isRecording && amplitude < sensitivity)
        {
            silenceTime += Time.deltaTime;
            Debug.Log("何も聞こえん");

        }

        // 録音を開始する閾値を超えた場合に録音を開始する
        if (!isRecording && amplitude > sensitivity)
        {
            recordedClip = Microphone.Start(null, true, (int)minSilenceTime, AudioSettings.outputSampleRate);
            isRecording = true;
            Debug.Log("Start Recording");
        }

        // 録音中に一定以上の無音時間がある場合に録音を終了する
        if (isRecording && silenceTime > minSilenceTime)
        {
            isRecording = false;
            //SavWav.Save("recordedClip.wav", recordedClip);
            Wav.ExportWav(recordedClip, outputFilePath);
            //Microphone.End(null);
            Debug.Log("End Recording");
            recordedClip = recordedClip = Microphone.Start(null, true, (int)minSilenceTime, AudioSettings.outputSampleRate);
        }
    }
        
    
}

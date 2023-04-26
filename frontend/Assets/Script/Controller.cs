using System;
using System.Linq;
using UnityEngine;

public class Controller : MonoBehaviour {
    
    [SerializeField] private string m_DeviceName;
    public float sensitivity = 0.1f; // 録音を開始する音量の閾値
    public float maxRecordTime = 10f; // 最大録音時間
    public float minSilenceTime = 1f; // 録音を終了するための最小の無音時間
    public string outputFilePath = "Assets/output.wav";

    private AudioClip m_AudioClip;
    private int m_LastAudioPos;
    private float m_AudioLevel;
    private bool isRecording = false;
    private float silenceTime = 0f;
    private float amplitude = 0;

    private string targetDevice = "";

    
    [SerializeField, Range(10, 100)] private float m_AmpGain = 10;
    
    void Start() {
        
        
        foreach (var device in Microphone.devices) {
            Debug.Log($"Device Name: {device}");
            if (device.Contains(m_DeviceName)) {
                targetDevice = device;
            }
        }
        
        Debug.Log($"=== Device Set: {targetDevice} ===");
        m_AudioClip = Microphone.Start(targetDevice, true, 10, 48000);
    }

    void Update() {
        float[] waveData = GetUpdatedAudio();
        if (waveData.Length == 0) return;
        
        m_AudioLevel = waveData.Average(Mathf.Abs);
        //m_Cube.transform.localScale = new Vector3(1, 1 + m_AmpGain * m_AudioLevel, 1);
        amplitude = m_AmpGain * m_AudioLevel;
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
            m_AudioClip = Microphone.Start(targetDevice, true, 10, 48000);
            isRecording = true;
            Debug.Log("Start Recording");
        }

        // 録音中に一定以上の無音時間がある場合に録音を終了する
        if (isRecording && silenceTime > minSilenceTime)
        {
            isRecording = false;
            //SavWav.Save("recordedClip.wav", recordedClip);
            Wav.ExportWav(m_AudioClip, outputFilePath);
            //Microphone.End(null);
            Debug.Log("End Recording");
            m_AudioClip = Microphone.Start(targetDevice, true, 10, 48000);
        }

    }
    
    private float[] GetUpdatedAudio() {
        
        int nowAudioPos = Microphone.GetPosition(null);// nullでデフォルトデバイス
        
        float[] waveData = Array.Empty<float>();

        if (m_LastAudioPos < nowAudioPos) {
            int audioCount = nowAudioPos - m_LastAudioPos;
            waveData = new float[audioCount];
            m_AudioClip.GetData(waveData, m_LastAudioPos);
        } else if (m_LastAudioPos > nowAudioPos) {
            int audioBuffer = m_AudioClip.samples * m_AudioClip.channels;
            int audioCount = audioBuffer - m_LastAudioPos;
            
            float[] wave1 = new float[audioCount];
            m_AudioClip.GetData(wave1, m_LastAudioPos);
            
            float[] wave2 = new float[nowAudioPos];
            if (nowAudioPos != 0) {
                m_AudioClip.GetData(wave2, 0);
            }

            waveData = new float[audioCount + nowAudioPos];
            wave1.CopyTo(waveData, 0);
            wave2.CopyTo(waveData, audioCount);
        }

        m_LastAudioPos = nowAudioPos;

        return waveData;
    }
}

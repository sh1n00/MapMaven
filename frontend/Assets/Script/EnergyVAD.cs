using UnityEngine;
using System.Collections;
using System.Collections.Generic;
using static Wav;

public class EnergyVAD : MonoBehaviour
{
    public AudioSource audioSource;    // 録音するオーディオソース
    public float sensitivity = 0.01f;  // 閾値

    private bool isRecording = false;  // 録音フラグ
    private AudioClip recordedClip;    // 録音されたクリップ
    private float minSilenceTime = 0.5f; // 無音時間の閾値
    private float silenceTime = 0.0f;  // 無音時間

    void Update()
    {
        // 音量を計算する
        float[] samples = new float[1024];
        audioSource.clip.GetData(samples, 0);
        float sum = 0.0f;
        for (int i = 0; i < samples.Length; i++)
        {
            sum += Mathf.Abs(samples[i]);
        }
        float amplitude = sum / samples.Length;

        // 録音中に音量が閾値以下になった場合、無音時間を加算する
        if (isRecording && amplitude < sensitivity)
        {
            silenceTime += Time.deltaTime;
        }

        // 録音を開始する閾値を超えた場合に録音を開始する
        if (!isRecording && amplitude > sensitivity)
        {
            recordedClip = Microphone.Start(null, true, (int)minSilenceTime, AudioSettings.outputSampleRate);
            isRecording = true;
        }

        // 録音中に一定以上の無音時間がある場合に録音を終了する
        if (isRecording && silenceTime > minSilenceTime)
        {
            isRecording = false;
            //SavWav.Save("recordedClip.wav", recordedClip);
            Wav.ExportWav(recordedClip,"Assets/output.wav");
            Microphone.End(null);
        }
    }
}


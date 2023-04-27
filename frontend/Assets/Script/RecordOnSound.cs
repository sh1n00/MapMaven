using UnityEngine;
using UnityEngine.Audio;

public class RecordOnSound : MonoBehaviour
{
    public AudioSource audioSource;
    public float threshold = 0.1f;
    public int sampleSize = 1024;
    public int recordingDuration = 5;
    public int bufferLength = 10;

    private AudioClip recordingClip;
    private float[] samples;
    private int bufferCount;

    void Start()
    {
        samples = new float[sampleSize];
        bufferCount = bufferLength;
        StartRecording();
    }

    void Update()
    {
        audioSource.GetOutputData(samples, 0);

        float sum = 0f;
        for (int i = 0; i < sampleSize; i++)
        {
            sum += Mathf.Abs(samples[i]);
        }

        float rmsValue = Mathf.Sqrt(sum / sampleSize);
        Debug.Log(rmsValue);
        if (rmsValue >= threshold && recordingClip == null)
        {
            StartRecording();
            Debug.Log("Start Record");
        }

        if (recordingClip != null)
        {
            recordingDuration--;
            if (recordingDuration <= 0)
            {
                StopRecording();
                Debug.Log("StopRecording");
            }
        }
    }

    void StartRecording()
    {
        if (bufferCount <= 0)
        {
            recordingClip = Microphone.Start(null, false, recordingDuration, AudioSettings.outputSampleRate);
            bufferCount = bufferLength;
        }
        else
        {
            bufferCount--;
        }
    }

    void StopRecording()
    {
        Microphone.End(null);
        //SavWav.Save(Application.dataPath + "/recording.wav", recordingClip);
        Wav.ExportWav(recordingClip,"Assets/record.wav");
        Debug.Log("ExportWav");
        recordingClip = null;
    }
}

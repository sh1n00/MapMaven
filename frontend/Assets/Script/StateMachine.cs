using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using GoogleCloudStreamingSpeechToText;


public class StateMachine : MonoBehaviour
{
    public AudioSource audioSource;
    public AudioClip thinkingClip;
    private bool isAudioPlay;
    public StreamingRecord streamingRecord;
    private bool isStandBy = false;
    private bool isAwaitReply = false;
    private enum gameState
    {
        standby,
        recognition,
        awaitReply,
        reply
    }
    private gameState _currentState = gameState.standby;


    private void LateUpdate()
    {
        //Debug.Log(_currentState);
        switch(_currentState)
        {
            case gameState.standby://not user 
                if(!isStandBy)streamingRecord.StartListening();
                isStandBy = true;
                isAudioPlay = false;
                break;
            case gameState.recognition://user talk start
                isStandBy = false;
                break;
            case gameState.awaitReply://user talk end and chatgpt
                if (!isStandBy) streamingRecord.StartListening();
                if (!isAwaitReply)
                {
                    audioSource.PlayOneShot(thinkingClip);
                }
                isAwaitReply = true;

                break;
            case gameState.reply://avater talk
                isStandBy = false;
                isAwaitReply = false;
                if (!audioSource.isPlaying && isAudioPlay)
                {
                    _currentState = gameState.standby;
                    
                }
                break;
        }
    }
    public async void EndRecognition(string finalDetection)
    {
        _currentState = gameState.awaitReply;
        streamingRecord.StopListening();
        var text = await HttpTest.GuideByText(finalDetection);
        Debug.Log(text);
        var test = await HttpTest.TextToAudio("2", "true", text);
        byte[] binaryData = Convert.FromBase64String(test);
        AudioClip clip = Wav.ToAudioClip(binaryData, "test");
        audioSource.PlayOneShot(clip);
        _currentState = gameState.reply;
        isAudioPlay = true;
    }
    public void StartRecognition()
    {
        _currentState= gameState.recognition;

    }


}

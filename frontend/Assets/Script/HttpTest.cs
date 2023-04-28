using System;
using System.Collections;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;
using Types;
using UnityEngine;

public class HttpTest : MonoBehaviour
{
    
    //private HttpClient _httpClient = new HttpClient();

    async void Start()
    {
        var test = await Chat("Hello");
        Debug.Log(test);
        //byte[] data = System.Text.Encoding.UTF8.GetBytes(test);
        //float[] samples = new float[data.Length / 4]; //size of a float is 4 bytes
        //Buffer.BlockCopy(data, 0, samples, 0, data.Length);
     
        //int channels = 1; //Assuming audio is mono because microphone input usually is
        //int sampleRate = 27000; //Assuming your samplerate is 44100 or change to 48000 or whatever is appropriate
        //Debug.Log(sampleRate);
        //AudioClip clip = AudioClip.Create("ClipName", samples.Length, channels, sampleRate, false);
        //clip.SetData(samples, 0);
        ////
        //// // Play the AudioClip
        //AudioSource.PlayClipAtPoint(clip, transform.position);
    }

    private async Task<String> TextToAudio(string speaker, string enable_interrogative_upspeak, string text)
    {

        HttpClient _httpClient = new HttpClient();
        var url = "http://localhost:8080/text-to-audio";
        var query = new Dictionary<String, String>
        {
            {"speaker", speaker},
            {"enable_interrogative_upspeak", enable_interrogative_upspeak},
            { "text", text },
        };
        var queryString = System.Web.HttpUtility.ParseQueryString("");
     
        foreach (KeyValuePair<String, String> pair in query)
        {
            queryString.Add(pair.Key, pair.Value);   
        }
       
        var uriBuilder = new UriBuilder(url)
        {
            Query = queryString.ToString()
        };
        
        try
        {

            var response = await _httpClient.GetAsync(uriBuilder.Uri);
            var responseBody = await response.Content.ReadAsStringAsync();
            Debug.Log(responseBody);
            AudioResponse json = JsonUtility.FromJson<AudioResponse>(responseBody);
            return json.AudioBinary;
        }
        catch (HttpRequestException e)
        {
            Debug.Log($"HTTP request failed: {e}");
            return "Error";
        }
    }
    
    public static async Task<string> Chat(string text)
    {
        HttpClient _httpClient = new HttpClient();
        var url = "http://localhost:8080/chat";
        var query = new Dictionary<String, String>
        {
            { "content", text }
        };
        var queryString = System.Web.HttpUtility.ParseQueryString("");
     
        foreach (KeyValuePair<String, String> pair in query)
        {
            queryString.Add(pair.Key, pair.Value);   
        }
       
        var uriBuilder = new UriBuilder(url)
        {
            Query = queryString.ToString()
        };
        
        try
        {
            var response = await _httpClient.GetAsync(uriBuilder.Uri);
            var responseBody = await response.Content.ReadAsStringAsync();
            ChatCompletionResponse json = JsonUtility.FromJson<ChatCompletionResponse>(responseBody);
            return json.choices[0].message.content;
        }
        catch (HttpRequestException e)
        {
            Debug.Log($"HTTP request failed: {e}");
            return "Error";
        }
    }

    // Update is called once per frame
    void Update()
    {
       
    }
}

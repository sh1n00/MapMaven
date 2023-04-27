using System;
using System.Collections;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;
using Types;
using UnityEngine;

public class HttpTest : MonoBehaviour
{
    private HttpClient _httpClient = new HttpClient();

    async void Start()
    {
        AudioSource audioSource = GetComponent<AudioSource>();
        var test1 = await GuideByText("サイバネティクスの場所はどこですか?");
        var test = await TextToAudio("1", "true", test1);
        byte[] binaryData = Convert.FromBase64String(test);
        AudioClip clip = Wav.ToAudioClip(binaryData, "test");
        audioSource.PlayOneShot(clip);
    }

    private async Task<string> GuideByText(string text)
    {
        var url = "http://localhost:8080/guide-by-text";
        var query = new Dictionary<String, String>
        {
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
            ChatCompletionResponse json = JsonUtility.FromJson<ChatCompletionResponse>(responseBody);
            return json.choices[0].message.content;
        }
        catch (HttpRequestException e)
        {
            Debug.Log($"HTTP request failed: {e}");
            return "text";
        }
    }

    private async Task<string> TextToAudio(string speaker, string enable_interrogative_upspeak, string text)
    {
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
            AudioResponse json = JsonUtility.FromJson<AudioResponse>(responseBody);
            return json.audio_binary;
        }
        catch (HttpRequestException e)
        {
            Debug.Log($"HTTP request failed: {e}");
            return "text";
        }
    }
    
    private async Task<string> Chat(string text)
    {
        var url = "http://localhost:8080/chat";
        var query = new Dictionary<String, String>
        {
            { "text", text }
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

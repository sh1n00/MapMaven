using System;
using System.Collections;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;
using Types;
using UnityEngine;

public class ChatTest : MonoBehaviour
{
    private HttpClient _httpClient;
    // Start is called before the first frame update
    async void Start()
    {
        var test = await Chat("Hello");
    }

    // Update is called once per frame
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
}

using System;
using System.Collections;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using UnityEngine;
using UnityEngine.Networking;

public class HttpRequest
{
    private HttpClient _httpClient;
    private HttpRequestMessage _request;
    public HttpRequest(HttpMethod method)
    {
        
        _request = new HttpRequestMessage{Method=method};
    }

    // ReSharper disable Unity.PerformanceAnalysis
    public async Task Send(String url, Dictionary<String, String> query)
    {
        var queryString = System.Web.HttpUtility.ParseQueryString("");
        if (query != null)
        {
            foreach (KeyValuePair<String, String> pair in query)
            {
                queryString.Add(pair.Key, pair.Value);   
            }
        }
        var uriBuilder = new System.UriBuilder(url)
        {
            Query = queryString.ToString()
        };

        _request.RequestUri = uriBuilder.Uri;
        HttpResponseMessage response =  await Task.Run(() => _httpClient.SendAsync(_request));
    }
}

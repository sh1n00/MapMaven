using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Networking;

public class HttpRequestTest : MonoBehaviour
{
    public string dataUrl;
    // Start is called before the first frame update
    void Start()
    {
        StartCoroutine(GetData());
    }

    // Update is called once per frame
    void Update()
    {
        
    }
    IEnumerator GetData()
    {
        //UnityWebRequestのインスタンスを作成します。
        UnityWebRequest req = UnityWebRequest.Get(dataUrl);

        // //リクエストに必要なヘッダーを設定します。Content-Typeヘッダーには、送信するデータの形式を指定します。
        // req.SetRequestHeader("Content-Type", "application/octet-stream");
        // //リクエストにリクエストに送信するデータを設定
        // req.uploadHandler = new UploadHandlerRaw(AudioClipGenerate.wavBase64);
        // req.uploadHandler.contentType = "application/octet-stream";


        yield return req.SendWebRequest();

        if (req.isNetworkError || req.isHttpError)
        {
            Debug.Log(req.error);
        }
        else if (req.responseCode == 200)
        {
            Debug.Log(req.downloadHandler.text);
        }
    }
}

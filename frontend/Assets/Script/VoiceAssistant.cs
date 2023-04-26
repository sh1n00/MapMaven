using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Windows.Speech;
using System.Collections.Generic;
using System.Linq;

public class VoiceAssistant : MonoBehaviour
{
    KeywordRecognizer keywordRecognizer;
    Dictionary<string, System.Action> keywords = new Dictionary<string, System.Action>();
    // Start is called before the first frame update
    void Start()
    {
        //Create keywords for keyword recognizer
        keywords.Add("activate", () =>
        {
            Debug.Log("activate");
        });

        keywordRecognizer = new KeywordRecognizer(keywords.Keys.ToArray());
        

        keywordRecognizer.Start();
        
    }

    private void KeywordRecognizer_OnPhraseRecognized(PhraseRecognizedEventArgs args)
    {
        System.Action keywordAction;
        // if the keyword recognized is in our dictionary, call that Action.
        if (keywords.TryGetValue(args.text, out keywordAction))
        {
            keywordAction.Invoke();
        }

    }
    

    // Update is called once per frame
    void Update()
    {
        keywordRecognizer.OnPhraseRecognized += KeywordRecognizer_OnPhraseRecognized;
    }
}

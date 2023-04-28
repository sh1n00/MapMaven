using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class FaceController : MonoBehaviour
{
    public SkinnedMeshRenderer _skinMesh;
    // Start is called before the first frame update
    void Start()
    {
        InvokeRepeating(nameof(RandomBlink), 2.0f, 3.0f);
        
    }

    // Update is called once per frame
    void Update()
    {
        
        
    }
    private void OnDestroy()
    {
        // DestroyŽž‚É“o˜^‚µ‚½Invoke‚ð‚·‚×‚ÄƒLƒƒƒ“ƒZƒ‹
        CancelInvoke();
    }
    private void RandomBlink()
    {
        if (Random.value > 0.3f)
        {
            _skinMesh.SetBlendShapeWeight(13, 100);
            Invoke("OpenEye", 0.3f);
            
        }
    }
    private void OpenEye()
    {
        _skinMesh.SetBlendShapeWeight(13, 0);
    }
}

using System;
using UnityEngine.Serialization;

namespace Types
{
    [Serializable]
    public class ChatCompletionResponse
    {
        public string id;
        public string @object;
        public long created;
        public Choice[] choices;
        public Usage usage;

        [Serializable]
        public class Choice
        {
            public int index;
            public Message message;
            [FormerlySerializedAs("finish_reason")] public string finishReason;

            [Serializable]
            public class Message
            {
                public string role;
                public string content;
            }
        }

        [Serializable]
        public class Usage
        {
            [FormerlySerializedAs("prompt_tokens")] public int promptTokens;
            [FormerlySerializedAs("completion_tokens")] public int completionTokens;
            [FormerlySerializedAs("total_tokens")] public int totalTokens;
        }
    }
}
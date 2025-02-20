using System.Text.Json.Serialization;

namespace GenLib;

public class Response
{
    
    [JsonPropertyName("message")]
    
    public string Message { get; set; }
    
}
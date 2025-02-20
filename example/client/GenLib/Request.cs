using System.Text.Json.Serialization;

namespace GenLib;

public class Request
{
    
    [JsonIgnore]
    [PathPropertyName("name")]
    
    public string Name { get; set; }
    
}
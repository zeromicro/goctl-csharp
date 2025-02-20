namespace GenLib;

public sealed class DemoApiClient : ApiBaseClient
{
    public DemoApiClient(string host, short port, string scheme = "http") : base(host, port, scheme){}
    
    public async Task<Response> GetFromNameAsync(Request request, CancellationToken cancellationToken, HttpContent? body=null)
    {
        return await RequestResultAsync<Request,Response>(
            HttpMethod.Get,
            "/from/:name",
            request,
            cancellationToken,
            body
        );
    }
}
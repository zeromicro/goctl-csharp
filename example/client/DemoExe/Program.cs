using GenLib;

try
{
    var client = new DemoApiClient("127.0.0.1", 8888);
    var request = new Request { Name = "me" }; // "you" or "me"
    var response = await client.GetFromNameAsync(request, CancellationToken.None);
    Console.WriteLine(response.Message);
}
catch (Exception e)
{
    Console.WriteLine(e.Message);
}
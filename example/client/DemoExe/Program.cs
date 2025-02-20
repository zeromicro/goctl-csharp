using GenLib;

try
{
    var client = new DemoApiClient("127.0.0.1", 8888);
    var request = new Request { Name = "me" }; // "you" or "me"
    var response = await client.GetFromNameAsync(request, CancellationToken.None);
    Console.WriteLine(response.Message);

    // will throw exception
    var request2 = new Request { Name = "me1" }; // options= "you" or "me"
    var response2 = await client.GetFromNameAsync(request2, CancellationToken.None);
}
catch (ApiException e)
{
    Console.WriteLine(e.Message);
}
catch (Exception e)
{
    Console.WriteLine(e.Message);
}
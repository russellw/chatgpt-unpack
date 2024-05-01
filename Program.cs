using System.Text.Json;

class Program {
	static void Main(string[] args) {
		// Path to the JSON file containing the ChatGPT messages.
		string filePath = "conversations.json";

		// Read the entire file into a string.
		string jsonContent = File.ReadAllText(filePath);

		// Deserialize the JSON content into a List of Conversation objects.
		List<Conversation> conversations = JsonSerializer.Deserialize<List<Conversation>>(jsonContent);

		// Reverse the order of conversations to start from the most recent.
		conversations.Reverse();

		// Loop through each conversation and its messages.
		foreach (var conversation in conversations) {
			var update_time = DateTimeOffset.FromUnixTimeSeconds((long)conversation.update_time).DateTime;
			Console.Write(update_time.ToString("yyyy-MM-dd"));
			Console.Write(":: ");
			Console.WriteLine(conversation.title);

			// Loop through each message in the mapping and print its content.
			if (conversation.mapping != null)
				foreach (var messageEntry in conversation.mapping) {
					var message = messageEntry.Value.Message;
					if (message != null && message.Content != null && message.Content.Parts != null) {
						Console.WriteLine("Message Content: " + string.Join("\n", message.Content.Parts));
					}
				}

			Console.WriteLine(); // Print a blank line for better readability.
		}
	}
}

public class Conversation {
	public string title { get; set; }
	public double update_time { get; set; }
	public Dictionary<string, MessageNode> mapping { get; set; }
}

public class MessageNode {
	public string Id { get; set; }
	public Message Message { get; set; }
}

public class Message {
	public string Id { get; set; }
	public Author Author { get; set; }
	public double? CreateTime { get; set; }
	public Content Content { get; set; }
}

public class Author {
	public string Role { get; set; }
	public string Name { get; set; }
}

public class Content {
	public string ContentType { get; set; }
	public List<string> Parts { get; set; }
}

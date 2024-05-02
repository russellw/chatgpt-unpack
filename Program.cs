using System.Text.Json;

static class Program {
	static void Main() {
		// Path to the JSON file containing the ChatGPT messages.
		string filePath = "conversations.json";

		// Read the entire file into a string.
		string jsonContent = File.ReadAllText(filePath);

		// Deserialize the JSON content into a List of Conversation objects.
		List<Conversation> conversations = JsonSerializer.Deserialize<List<Conversation>>(jsonContent);

		// Reverse the order of conversations to start from the most recent.
		conversations.Reverse();

		using StreamWriter writer = new("conversations.txt");

		// Loop through each conversation and its messages.
		foreach (var conversation in conversations) {
			var update_time = DateTimeOffset.FromUnixTimeSeconds((long)conversation.update_time).DateTime;
			writer.Write(update_time.ToString("yyyy-MM-dd"));
			writer.Write(":: ");
			writer.WriteLine(conversation.title);

			// Loop through each message in the mapping and print its content.
			if (conversation.mapping != null)
				foreach (var messageEntry in conversation.mapping) {
					var message = messageEntry.Value.message;
					var text = Text(message);
					if (string.IsNullOrEmpty(text))
						continue;
					var user = User(message);
					writer.Write(user ? "Q: " : "A: ");
					writer.WriteLine(text);
					if (!user)
						writer.WriteLine(); // Print a blank line for better readability.
				}
		}
	}

	static string Text(Message message) {
		if (message == null)
			return null;
		if (message.content == null)
			return null;
		if (message.content.parts == null)
			return null;
		return string.Join("\n", message.content.parts);
	}

	static string Role(Message message) {
		if (message == null)
			return null;
		if (message.author == null)
			return null;
		return message.author.role;
	}

	static bool User(Message message) {
		return Role(message) == "user";
	}
}

class Conversation {
	public string title { get; set; }
	public double update_time { get; set; }
	public Dictionary<string, MessageNode> mapping { get; set; }
}

class MessageNode {
	public Message message { get; set; }
}

class Message {
	public Author author { get; set; }
	public Content content { get; set; }
}

class Author {
	public string role { get; set; }
}

class Content {
	public string ContentType { get; set; }
	public List<JsonElement> parts { get; set; }
}

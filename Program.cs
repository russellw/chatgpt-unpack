using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

namespace ChatGPTMessageReader
{
    class Program
    {
        static void Main(string[] args)
        {
            // Path to the JSON file containing the ChatGPT messages.
            string filePath = "conversations.json";

            // Read the entire file into a string.
            string jsonContent = File.ReadAllText(filePath);

            // Deserialize the JSON content into a List of Conversation objects.
            List<Conversation> conversations = JsonSerializer.Deserialize<List<Conversation>>(jsonContent);

            // Loop through each conversation and its messages.
            foreach (var conversation in conversations)
            {
                Console.WriteLine("Title: " + conversation.Title);
                Console.WriteLine("Created On: " + DateTimeOffset.FromUnixTimeSeconds((long)conversation.CreateTime).DateTime);

                // Loop through each message in the mapping and print its content.
                foreach (var messageEntry in conversation.Mapping)
                {
                    var message = messageEntry.Value.Message;
                    if (message != null && message.Content != null && message.Content.Parts != null)
                    {
                        Console.WriteLine("Message Content: " + string.Join("\n", message.Content.Parts));
                    }
                }

                Console.WriteLine(); // Print a blank line for better readability.
            }
        }
    }

    public class Conversation
    {
        public string Title { get; set; }
        public double CreateTime { get; set; }
        public Dictionary<string, MessageNode> Mapping { get; set; }
    }

    public class MessageNode
    {
        public string Id { get; set; }
        public Message Message { get; set; }
    }

    public class Message
    {
        public string Id { get; set; }
        public Author Author { get; set; }
        public double? CreateTime { get; set; }
        public Content Content { get; set; }
    }

    public class Author
    {
        public string Role { get; set; }
        public string Name { get; set; }
    }

    public class Content
    {
        public string ContentType { get; set; }
        public List<string> Parts { get; set; }
    }
}

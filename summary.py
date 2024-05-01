import json

def summarize_json_structure(data):
    if isinstance(data, dict):
        return {key: summarize_json_structure(value) for key, value in data.items()}
    elif isinstance(data, list):
        # Handle non-empty lists; generalize the first item (assuming homogeneity)
        if data:
            return [summarize_json_structure(data[0])]  # Adjust if list items vary
        return []
    else:
        # Return the type of the item
        return str(type(data).__name__)

# Load JSON data from a file
with open('conversations.json', 'r') as file:
    json_data = json.load(file)

# Summarize the JSON structure
json_structure = summarize_json_structure(json_data)

# Print the summarized structure with indentation for better readability
print(json.dumps(json_structure, indent=4))

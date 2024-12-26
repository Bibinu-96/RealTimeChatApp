from flask import Flask, request, jsonify
from ollama import Client


app = Flask(__name__)

@app.route("/process", methods=["POST"])
def process():
    try:
        # Get JSON payload
        data = request.get_json()
        if not data:
            return jsonify({"error": "No JSON payload provided"}), 400

        # Extract the input JSON string
        json_string = data.get("query", "")
        if not json_string:
            return jsonify({"error": "json_string is required"}), 400

        # Call the Ollama command with the input
        try:
            client = Client(
                host='http://localhost:11434',
                headers={'x-some-header': 'some-value'}
                )
            response = client.chat(model='smollm:latest', messages=[
                {
                    'role': 'user',
                    'content': json_string,
                    },
                ])
            result = response['message']['content']
            return jsonify({"result": result})
        except Exception as e:
            return jsonify({"error": f"Error running Ollama: {e.stderr}"}), 500

    except Exception as e:
        return jsonify({"error": str(e)}), 500


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)

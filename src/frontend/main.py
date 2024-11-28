from flask import Flask, jsonify, request

# Initialize the Flask application
app = Flask(__name__)

# Define a route for the root URL
@app.route('/')
def home():
    return "<h1>🚀 Welcome to Your Flask Server!</h1><p>Flask makes web development easy!</p>"

# Define a route that returns JSON data
@app.route('/api/data', methods=['GET'])
def get_data():
    # Example JSON response
    response_data = {
        "message": "Hello, Flask!",
        "status": "success",
        "data": {
            "framework": "Flask",
            "version": "2.x",
        }
    }
    return jsonify(response_data)

# Define a POST route for handling data
@app.route('/api/submit', methods=['POST'])
def submit_data():
    # Get data from the request body (assuming JSON)
    data = request.json
    if not data:
        return jsonify({"error": "No JSON body provided"}), 400
    return jsonify({"message": "Data received!", "data": data}), 200

# Start the server
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)


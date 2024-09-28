// custom_client.dart
import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:safe_haven/features/auth/data/data_sources/local_data_source.dart';

abstract class NetworkInfo {
  Future<bool> get isConnected;
}

class NetworkInfoImpl implements NetworkInfo {
  final InternetConnectionChecker connectionChecker;

  NetworkInfoImpl(this.connectionChecker);

  @override
  Future<bool> get isConnected => connectionChecker.hasConnection;
}

class CustomHttpClient {
  final http.Client _client;
  final AuthenticationLocalDataSource _authenticationLocalDataSource;

  CustomHttpClient({
    required http.Client client,
    required AuthenticationLocalDataSource authenticationLocalDataSource,
  })  : _client = client,
        _authenticationLocalDataSource = authenticationLocalDataSource;

  Future<http.Response> get(String endpoint) async {
    return _client.get(
      _parseUrl(endpoint),
      headers: await _headers(),
    );
  }

  Future<http.Response> post(String endpoint, {Object? body}) async {
    final jsonBody = body != null ? jsonEncode(body) : null;
    return _client.post(
      _parseUrl(endpoint),
      body: jsonBody,
      headers: await _headers(),
    );
  }

  Future<http.Response> put(String endpoint, {Object? body}) async {
    return _client.put(
      _parseUrl(endpoint),
      body: body,
      headers: await _headers(),
    );
  }

  Future<http.Response> delete(String endpoint) async {
    return _client.delete(
      _parseUrl(endpoint),
      headers: await _headers(),
    );
  }

  Future<http.StreamedResponse> send(http.BaseRequest request) async {
    request.headers.addAll(await _headers());
    return request.send();
  }

  Uri _parseUrl(String endpoint) {
    return Uri.parse(endpoint);
  }

  Future<Map<String, String>> _headers() async {
    final token = await _authenticationLocalDataSource.getToken();
    return {
      'Authorization': 'Bearer $token',
      'Content-Type': 'application/json',
    };
  }
}

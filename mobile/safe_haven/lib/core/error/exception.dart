
class CacheException implements Exception {}

class SocketException implements Exception {
  SocketException(String socketError);
}

class NotFoundException implements Exception {
  final String message;
  NotFoundException(this.message);

  @override
  String toString() => 'NotFoundException: $message';
}

class JsonParsingException implements Exception {}



class ImageException implements Exception {}

class UnauthorizedException implements Exception {}

class UserAlreadyExistsException implements Exception {}

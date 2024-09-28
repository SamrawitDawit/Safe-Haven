class Urls {
  static const String authUrl = 'http://192.168.1.10:8080/auth';
}

class ErrorMessages {
  static const String noInternet = 'Failed to connect to the internet';
  static const String somethingWentWrong = 'Something went wrong';
  static const String serverError = 'An error has occurred';
  static const String cacheError = 'Failed to load cache';
  static const String socketError =
      'No Internet connection or server unreachable';
  static const String forbiddenError = 'Invalid Credentials! Please try again';
  static const String userAlreadyExists = 'User Already Exists';
  static const String notFoundError = 'No such account';
}

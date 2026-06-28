import 'package:flutter/material.dart';
import 'package:go_blog_app/router/app_router.dart';
import 'package:go_blog_app/theme/app_theme.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'Go Blog',
      theme: AppTheme.light,
      routerConfig: router,
    );
  }
}

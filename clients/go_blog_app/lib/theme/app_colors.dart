import 'package:flutter/material.dart';

/// Color tokens from [internal/web/static/css/input.css].
abstract final class AppColors {
  static const Color surface = Color(0xFFFDF8F8);
  static const Color surfaceDim = Color(0xFFDDD9D8);
  static const Color surfaceBright = Color(0xFFFDF8F8);
  static const Color surfaceContainerLowest = Color(0xFFFFFFFF);
  static const Color surfaceContainerLow = Color(0xFFF7F3F2);
  static const Color surfaceContainer = Color(0xFFF1EDEC);
  static const Color surfaceContainerHigh = Color(0xFFEBE7E6);
  static const Color surfaceContainerHighest = Color(0xFFE5E2E1);
  static const Color onSurface = Color(0xFF1C1B1B);
  static const Color onSurfaceVariant = Color(0xFF444748);
  static const Color inverseSurface = Color(0xFF313030);
  static const Color inverseOnSurface = Color(0xFFF4F0EF);
  static const Color outline = Color(0xFF747878);
  static const Color outlineVariant = Color(0xFFC4C7C7);
  static const Color surfaceTint = Color(0xFF5F5E5E);
  static const Color primary = Color(0xFF000000);
  static const Color onPrimary = Color(0xFFFFFFFF);
  static const Color primaryContainer = Color(0xFF1C1B1B);
  static const Color onPrimaryContainer = Color(0xFF858383);
  static const Color inversePrimary = Color(0xFFC8C6C5);
  static const Color secondary = Color(0xFF4B41E1);
  static const Color onSecondary = Color(0xFFFFFFFF);
  static const Color secondaryContainer = Color(0xFF645EFB);
  static const Color onSecondaryContainer = Color(0xFFFFFBFF);
  static const Color tertiary = Color(0xFF000000);
  static const Color onTertiary = Color(0xFFFFFFFF);
  static const Color tertiaryContainer = Color(0xFF1C1B1A);
  static const Color onTertiaryContainer = Color(0xFF868382);
  static const Color error = Color(0xFFBA1A1A);
  static const Color onError = Color(0xFFFFFFFF);
  static const Color errorContainer = Color(0xFFFFDAD6);
  static const Color onErrorContainer = Color(0xFF93000A);
  static const Color primaryFixed = Color(0xFFE5E2E1);
  static const Color primaryFixedDim = Color(0xFFC8C6C5);
  static const Color onPrimaryFixed = Color(0xFF1C1B1B);
  static const Color onPrimaryFixedVariant = Color(0xFF474746);
  static const Color secondaryFixed = Color(0xFFE2DFFF);
  static const Color secondaryFixedDim = Color(0xFFC3C0FF);
  static const Color onSecondaryFixed = Color(0xFF0F0069);
  static const Color onSecondaryFixedVariant = Color(0xFF3323CC);
  static const Color tertiaryFixed = Color(0xFFE6E2DF);
  static const Color tertiaryFixedDim = Color(0xFFCAC6C4);
  static const Color onTertiaryFixed = Color(0xFF1C1B1A);
  static const Color onTertiaryFixedVariant = Color(0xFF484645);
  static const Color background = Color(0xFFFDF8F8);
  static const Color onBackground = Color(0xFF1C1B1B);
  static const Color surfaceVariant = Color(0xFFE5E2E1);

  static const ColorScheme colorScheme = ColorScheme(
    brightness: Brightness.light,
    primary: primary,
    onPrimary: onPrimary,
    primaryContainer: primaryContainer,
    onPrimaryContainer: onPrimaryContainer,
    secondary: secondary,
    onSecondary: onSecondary,
    secondaryContainer: secondaryContainer,
    onSecondaryContainer: onSecondaryContainer,
    tertiary: tertiary,
    onTertiary: onTertiary,
    tertiaryContainer: tertiaryContainer,
    onTertiaryContainer: onTertiaryContainer,
    error: error,
    onError: onError,
    errorContainer: errorContainer,
    onErrorContainer: onErrorContainer,
    surface: surface,
    onSurface: onSurface,
    onSurfaceVariant: onSurfaceVariant,
    outline: outline,
    outlineVariant: outlineVariant,
    shadow: Colors.black,
    scrim: Colors.black,
    inverseSurface: inverseSurface,
    onInverseSurface: inverseOnSurface,
    inversePrimary: inversePrimary,
    surfaceTint: surfaceTint,
    surfaceContainerHighest: surfaceContainerHighest,
  );
}

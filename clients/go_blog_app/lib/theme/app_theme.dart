import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

import 'app_colors.dart';
import 'app_spacing.dart';

/// Flutter theme derived from the web Minimalist Editorial design system.
abstract final class AppTheme {
  static ThemeData get light {
    final textTheme = _buildTextTheme();

    return ThemeData(
      useMaterial3: true,
      colorScheme: AppColors.colorScheme,
      scaffoldBackgroundColor: AppColors.background,
      textTheme: textTheme,
      appBarTheme: AppBarTheme(
        backgroundColor: AppColors.surface,
        foregroundColor: AppColors.onSurface,
        elevation: 0,
        scrolledUnderElevation: 0,
        surfaceTintColor: Colors.transparent,
        titleTextStyle: textTheme.headlineSmall,
      ),
      cardTheme: CardThemeData(
        color: AppColors.surfaceContainerLow,
        elevation: 0,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(AppSpacing.radiusMd),
          side: const BorderSide(color: AppColors.outlineVariant),
        ),
      ),
      dividerTheme: const DividerThemeData(
        color: AppColors.outlineVariant,
        thickness: 1,
        space: 1,
      ),
      filledButtonTheme: FilledButtonThemeData(
        style: FilledButton.styleFrom(
          backgroundColor: AppColors.primary,
          foregroundColor: AppColors.onPrimary,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(AppSpacing.radius),
          ),
          textStyle: textTheme.labelMedium,
        ),
      ),
      textButtonTheme: TextButtonThemeData(
        style: TextButton.styleFrom(
          foregroundColor: AppColors.secondary,
          textStyle: textTheme.labelMedium,
        ),
      ),
      inputDecorationTheme: InputDecorationTheme(
        filled: true,
        fillColor: AppColors.surfaceContainerLow,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(AppSpacing.radius),
          borderSide: const BorderSide(color: AppColors.outlineVariant),
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(AppSpacing.radius),
          borderSide: const BorderSide(color: AppColors.outlineVariant),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(AppSpacing.radius),
          borderSide: const BorderSide(color: AppColors.secondary, width: 2),
        ),
        errorBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(AppSpacing.radius),
          borderSide: const BorderSide(color: AppColors.error),
        ),
        labelStyle: textTheme.labelMedium?.copyWith(
          color: AppColors.onSurfaceVariant,
        ),
        hintStyle: textTheme.bodyMedium?.copyWith(
          color: AppColors.outline,
        ),
      ),
    );
  }

  static TextTheme _buildTextTheme() {
    final playfair = GoogleFonts.playfairDisplayTextTheme();
    final inter = GoogleFonts.interTextTheme();

    return TextTheme(
      displayLarge: playfair.displayLarge?.copyWith(
        fontSize: 48,
        fontWeight: FontWeight.w700,
        height: 1.2,
        letterSpacing: -0.02 * 48,
        color: AppColors.onSurface,
      ),
      displayMedium: playfair.displayMedium?.copyWith(
        fontSize: 36,
        fontWeight: FontWeight.w700,
        height: 1.2,
        color: AppColors.onSurface,
      ),
      headlineMedium: playfair.headlineMedium?.copyWith(
        fontSize: 32,
        fontWeight: FontWeight.w600,
        height: 1.3,
        color: AppColors.onSurface,
      ),
      headlineSmall: playfair.headlineSmall?.copyWith(
        fontSize: 24,
        fontWeight: FontWeight.w600,
        height: 1.4,
        color: AppColors.onSurface,
      ),
      bodyLarge: inter.bodyLarge?.copyWith(
        fontSize: 20,
        fontWeight: FontWeight.w400,
        height: 1.8,
        color: AppColors.onSurface,
      ),
      bodyMedium: inter.bodyMedium?.copyWith(
        fontSize: 16,
        fontWeight: FontWeight.w400,
        height: 1.6,
        color: AppColors.onSurface,
      ),
      labelMedium: inter.labelMedium?.copyWith(
        fontSize: 14,
        fontWeight: FontWeight.w500,
        height: 1.2,
        letterSpacing: 0.02 * 14,
        color: AppColors.onSurfaceVariant,
      ),
      labelSmall: inter.labelSmall?.copyWith(
        fontSize: 12,
        fontWeight: FontWeight.w600,
        height: 1,
        letterSpacing: 0.05 * 12,
        color: AppColors.onSurfaceVariant,
      ),
    );
  }
}

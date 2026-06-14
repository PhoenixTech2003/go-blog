---
name: Minimalist Editorial
colors:
  surface: '#fdf8f8'
  surface-dim: '#ddd9d8'
  surface-bright: '#fdf8f8'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#f7f3f2'
  surface-container: '#f1edec'
  surface-container-high: '#ebe7e6'
  surface-container-highest: '#e5e2e1'
  on-surface: '#1c1b1b'
  on-surface-variant: '#444748'
  inverse-surface: '#313030'
  inverse-on-surface: '#f4f0ef'
  outline: '#747878'
  outline-variant: '#c4c7c7'
  surface-tint: '#5f5e5e'
  primary: '#000000'
  on-primary: '#ffffff'
  primary-container: '#1c1b1b'
  on-primary-container: '#858383'
  inverse-primary: '#c8c6c5'
  secondary: '#4b41e1'
  on-secondary: '#ffffff'
  secondary-container: '#645efb'
  on-secondary-container: '#fffbff'
  tertiary: '#000000'
  on-tertiary: '#ffffff'
  tertiary-container: '#1c1b1a'
  on-tertiary-container: '#868382'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
  primary-fixed: '#e5e2e1'
  primary-fixed-dim: '#c8c6c5'
  on-primary-fixed: '#1c1b1b'
  on-primary-fixed-variant: '#474746'
  secondary-fixed: '#e2dfff'
  secondary-fixed-dim: '#c3c0ff'
  on-secondary-fixed: '#0f0069'
  on-secondary-fixed-variant: '#3323cc'
  tertiary-fixed: '#e6e2df'
  tertiary-fixed-dim: '#cac6c4'
  on-tertiary-fixed: '#1c1b1a'
  on-tertiary-fixed-variant: '#484645'
  background: '#fdf8f8'
  on-background: '#1c1b1b'
  surface-variant: '#e5e2e1'
typography:
  display-lg:
    fontFamily: Playfair Display
    fontSize: 48px
    fontWeight: '700'
    lineHeight: '1.2'
    letterSpacing: -0.02em
  display-lg-mobile:
    fontFamily: Playfair Display
    fontSize: 36px
    fontWeight: '700'
    lineHeight: '1.2'
  headline-md:
    fontFamily: Playfair Display
    fontSize: 32px
    fontWeight: '600'
    lineHeight: '1.3'
  headline-sm:
    fontFamily: Playfair Display
    fontSize: 24px
    fontWeight: '600'
    lineHeight: '1.4'
  body-lg:
    fontFamily: Inter
    fontSize: 20px
    fontWeight: '400'
    lineHeight: '1.8'
  body-md:
    fontFamily: Inter
    fontSize: 16px
    fontWeight: '400'
    lineHeight: '1.6'
  label-md:
    fontFamily: Inter
    fontSize: 14px
    fontWeight: '500'
    lineHeight: '1.2'
    letterSpacing: 0.02em
  label-sm:
    fontFamily: Inter
    fontSize: 12px
    fontWeight: '600'
    lineHeight: '1'
    letterSpacing: 0.05em
rounded:
  sm: 0.25rem
  DEFAULT: 0.5rem
  md: 0.75rem
  lg: 1rem
  xl: 1.5rem
  full: 9999px
spacing:
  unit: 8px
  container-max: 1120px
  reading-max: 720px
  gutter: 24px
  margin-mobile: 16px
  section-gap: 80px
---

## Brand & Style
The design system focuses on the "Minimalist Editorial" aesthetic, prioritizing legibility and content hierarchy over decorative elements. It targets writers, thinkers, and avid readers who value a distraction-free environment.

The UI leverages heavy whitespace and a restricted color palette to evoke a sense of calm authority and intellectual clarity. The emotional response should be one of "quiet focus"—a digital sanctuary for long-form thought. Drawing from high-end print journalism, the system uses "Corporate Modern" structural principles mixed with "Minimalist" restraint to ensure the interface recedes, allowing the author's words to take center stage.

## Colors
The palette is designed to reduce optical fatigue during long reading sessions. 
- **Primary (#1A1A1A):** Used for maximum contrast in typography and core branding elements.
- **Secondary (#4F46E5):** A muted indigo reserved for functional affordances like links, active states, and primary CTA buttons.
- **Background (#FAFAFA):** A subtle off-white that softens the "vibration" of black text on a pure white screen.
- **Surface (#FFFFFF):** Used for elevated elements like cards, modal overlays, and input fields to create subtle depth.
- **Functional Colors:** Soft Emerald and Soft Rose are used sparingly for status indicators to maintain the monochromatic integrity of the workspace.

## Typography
The typographic strategy follows a classic serif/sans-serif pairing. **Playfair Display** provides an editorial, literary character for headlines. **Inter** is used for body text and UI labels due to its exceptional legibility and neutral tone.

Key rules:
- **Body Text:** Always maintain a line height of at least 1.6 to 1.8 for long-form content. 
- **Column Width:** Limit reading widths to a maximum of 720px to prevent eye strain.
- **Hierarchy:** Use the Display-LG style exclusively for article titles.
- **Labels:** Use uppercase for small labels to distinguish functional UI from narrative content.

## Layout & Spacing
The layout follows a **Fixed Grid** approach for the main feed and a **No Grid (Centered)** approach for the reading experience.

- **Desktop:** A 12-column grid with a max width of 1120px for the discovery/feed pages.
- **Reading Mode:** Content is centered with a max-width of 720px, surrounded by generous 80px+ vertical gaps between sections (Title, Body, Comments).
- **Rhythm:** All spacing (margins, padding, gaps) should be multiples of the 8px base unit.
- **Mobile:** Margins scale down to 16px, and typography shifts to mobile-specific tokens to ensure high contrast and readability on small screens.

## Elevation & Depth
This design system avoids heavy shadows. Hierarchy is established through **Tonal Layers** and **Low-Contrast Outlines**.

- **Level 0 (Background):** #FAFAFA.
- **Level 1 (Surface/Cards):** #FFFFFF with a 1px border of #E5E5E5 (or 5% black). 
- **Depth:** Use very soft, diffused ambient shadows (blur 20px, 4% opacity) only for floating elements like dropdowns or navigation bars. 
- **Interaction:** On hover, cards may lift slightly with a subtle increase in shadow diffusion, but the border remains the primary anchor.

## Shapes
The shape language is "Rounded" (8px base) to soften the professional tone and make the platform feel modern and approachable. 

- **Buttons & Inputs:** Use the 8px (0.5rem) standard radius.
- **Cards:** Use 12px or 16px (1rem) for larger container elements.
- **Avatars:** Always circular (full-round) to contrast against the structured grid.

## Components
- **Buttons:** Primary buttons are #1A1A1A with white text. Secondary actions use #4F46E5 as a text-only link or a ghost button.
- **Article Cards:** Minimalist layouts. Image at the top (if present), followed by a Serif headline and a brief Sans-Serif excerpt. Avoid excessive metadata; show only "Date" and "Read Time."
- **Navigation:** A sticky top bar with #FAFAFA background and a thin bottom border. Keep links limited to 3-4 key areas to maintain focus.
- **Inputs:** Pure white background (#FFFFFF) with a 1px #E5E5E5 border. Focus state should use a 2px #4F46E5 border.
- **Chips/Tags:** Small, #F3F4F6 background with #1A1A1A text. No borders.
- **Reading Progress:** A thin 2px bar at the very top of the viewport in #4F46E5 to provide non-intrusive feedback to the reader.
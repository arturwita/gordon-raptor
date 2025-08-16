# Gordon Raptor

## AI Model Strategy

- **Primary:** Use `gpt-4o-mini` with real-time API for all AI interactions.
- **Fallback:** If costs are too high, switch to a voice-to-text model and send voice input as a text prompt.
- **Real-Time Function Calling:** Enable real-time function calling for dynamic interactions.

---

## User Experience & Flow

- **Web Client:**
    - Simple UI with a button to start/stop the conversation.
    - Voice input to choose and interact with recipes.
    - Pause conversation if silent for 1 minute or on explicit stop instruction.
- **Voice Assistance:**
    - Guides user through recipe steps.
    - Can answer questions and provide tips during cooking.
- **Recipe Selection:**
    - Users can choose a recipe by voice (MCP tool call).
    - List of available recipes shown in UI.
- **Timers:**
    - Multiple timers per session (Phase 2).
    - Timers managed via WebSocket, server-sent events, and persisted in Redis (with TTL).
- **Recipe Management:**
    - Start with a hardcoded recipe list (to be improved in later phases).
    - Option to improve/modify recipes (Phase 2).
    - Share/export recipes (Phase 3).
    - Videos for particular steps (Phase 3).
    - Suggest recipes based on fridge contents, with AI-powered ingredient recognition and substitutions (Phase 3).

---

## Features by Phase

### Phase 1: Core MVP
- Hardcoded recipe list.
- Voice-driven recipe selection and step-by-step guidance.
- Simple start/stop button.
- Real-time AI assistance.

### Phase 2: Enhanced Cooking Experience
- Multiple timers per session.
- Recipe improvement and editing.
- Persistent timers (Redis).
- UI/UX improvements (Stitch AI design).

### Phase 3: Advanced Features
- Recipe sharing/exporting.
- Step videos.
- Smart recipe suggestions based on fridge inventory (AI product recognition, substitutions).

---

## Technical Stack

- **AI:**
    - `gpt-4o-mini` (primary), fallback to voice-to-text + text prompt.
- **Backend:**
    - Go, Gin framework.
    - MongoDB (for recipe notes).
    - Redis (for timers, with TTL).
- **Frontend:**
    - React, Tailwind CSS.
    - Stitch AI for UI design.
- **Testing:**
    - Go for backend tests.
    - Node.js for end-to-end (e2e) tests.

---

## Context Storage: What Should Be Stored?

Besides the recipe itself, the following should be considered for context storage:

- **Current Recipe State:**
    - Which recipe is active.
    - Current step in the recipe.
    - Step history (for "go back" or "repeat" commands).
- **Timers:**
    - All active timers (label, duration, start time, remaining time).
    - Timer state (running, paused, finished).
- **User Preferences:**
    - Dietary restrictions, favorite cuisines, disliked ingredients.
    - Voice interaction preferences (speed, language, etc.).
- **Session Metadata:**
    - Session start time, last interaction time (for auto-pause).
    - User ID/session ID (for persistence).
- **Fridge/Inventory State:**
    - List of available ingredients (for smart suggestions).
- **Recipe Modifications:**
    - User notes, substitutions, or improvements to recipes.
- **Device/Environment Info:**
    - For tailoring instructions (e.g., oven type, measurement units).
- **Interaction History:**
    - Previous questions/answers, clarifications, and corrections.

---

## Next Steps

1. **MVP Implementation:**
    - Set up hardcoded recipes, voice input, and real-time AI guidance.
2. **Design UI:**
    - Use Stitch AI for a clean, accessible interface.
3. **Backend Setup:**
    - Implement Go/Gin server, MongoDB for notes, Redis for timers.
4. **Plan for Phased Rollout:**
    - Prepare for incremental feature releases as outlined above.

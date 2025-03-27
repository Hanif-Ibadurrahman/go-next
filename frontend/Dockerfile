# Build Stage
FROM node:20-alpine AS builder
WORKDIR /app

# Install build dependencies (needed for native modules)
RUN apk add --no-cache python3 make g++ libc6-compat

# Install dependencies with caching
COPY package.json package-lock.json* ./
# Remove --omit=dev to ensure all dependencies are installed for build
RUN npm ci

# Copy source code
COPY . .

# Build the Next.js application
RUN npm run build

# Prune dev dependencies after build
RUN npm prune --production

# Production Stage
FROM node:20-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production
ENV PORT=3000

# Install runtime dependencies
RUN apk add --no-cache libc6-compat

# Check if group/user exists before creating
RUN if ! getent group nodejs > /dev/null; then \
        addgroup -g 10001 nodejs; \
    fi && \
    if ! getent passwd nextjs > /dev/null; then \
        adduser -D -u 10001 -G nodejs nextjs; \
    fi

# Copy only necessary files with proper ownership
COPY --from=builder --chown=nextjs:nodejs /app/.next ./.next
COPY --from=builder --chown=nextjs:nodejs /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/node_modules ./node_modules
COPY --from=builder --chown=nextjs:nodejs /app/package.json ./package.json
COPY --from=builder --chown=nextjs:nodejs /app/next.config.js* ./

USER nextjs
EXPOSE 3000
CMD ["npm", "start"]
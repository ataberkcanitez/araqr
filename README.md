# araQR

## Getting Started

Build the application binary by running the following command:

```bash
make bin
```

Run the dependencies by running the following command:

```bash
docker-compose up -d
```

Run the following command to run the application in local:

```bash
go run . serve \
    --log.level debug \
    --db.sslmode disable \
    --auth.secret-key secret
```

**araQR** is a lightweight, privacy-aware platform that enables car owners to receive predefined messages from third
parties through a QR code sticker placed inside their vehicle. It helps people notify drivers (e.g. in case of improper
parking) without revealing personal contact information.

---

## 🚘 Why araQR?

In urban areas, especially densely populated cities, it's often difficult to reach vehicle owners when needed. Common
issues include:

- Improper parking
- Blocking exits or driveways
- Leaving lights on
- Forgotten items

araQR solves this by offering a one-way, privacy-preserving communication channel. Car owners control what information
is visible and are notified via mobile app when someone scans their QR sticker and selects a predefined message.

---


![Vertical Sequence](/assets/vertical_sequence.png)

![Sequence Diagram for Sticker Owner](/assets/sticker-owner-flow.png)
![3rd Party User Flow](/assets/3rd_party_user_flow.png)

## 📱 How It Works

### For Car Owners (QR Code Holders):

1. Purchase a unique QR code sticker.
2. Download the mobile app (React Native).
3. Scan and link the QR code to their account.
4. Customize profile visibility (e.g., license plate, car color).
5. Get real-time push notifications when someone sends a message

### For 3rd Parties (Message Senders):

1. Scan the QR code using their phone camera (no app required).
2. They are redirected to a browser page.
3. Choose from predefined messages (e.g., “You're blocking the garage.”) or write your message.
4. Message is instantly delivered as a push notification to the car owner's app.

---

## 📦 Technical Stack

| Layer            | Technology                |
|------------------|---------------------------|
| Mobile App       | React Native (Expo)       |
| Backend          | Golang                    |
| Hosting          | AWS App Runner            |
| Database         | postgresql                |
| Notifications    | ???                       |
| Domain & Routing | AWS Route 53 + CloudFront |
| CI/CD            | GitHub Actions            |

---
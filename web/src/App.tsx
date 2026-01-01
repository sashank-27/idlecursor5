import { useEffect, useMemo, useState } from "react";
import { AgentLogEntry, AgentStatus, getLogs, getStatus, openStatusStream, startSession, stopSession } from "./api";

export default function App() {
  const [status, setStatus] = useState<AgentStatus | null>(null);
  const [logs, setLogs] = useState<AgentLogEntry[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const stateLabel = useMemo(() => {
    if (!status) return "Unknown";
    if (status.policyLocked) return "Locked";
    return status.state === "active" ? "Active" : status.state === "paused" ? "Paused" : "Idle";
  }, [status]);

  useEffect(() => {
    refresh();

    // subscribe to live status updates (SSE)
    const unsubscribe = openStatusStream((s: AgentStatus) => {
      setStatus(s);
    });

    return () => {
      unsubscribe();
    };
  }, []);

  async function refresh() {
    setError(null);
    try {
      const [s, l] = await Promise.all([getStatus(), getLogs()]);
      setStatus(s);
      setLogs((l.entries || []).slice(-50));
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err));
    }
  }

  async function onStart() {
    setLoading(true);
    setError(null);
    try {
      // Use simple defaults: moderate randomness, 2s idle threshold to match request, long duration.
      await startSession({ mode: "meeting", randomness: 0.5, idleThresholdSeconds: 2, maxDurationMinutes: 240 });
      await refresh();
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err));
    } finally {
      setLoading(false);
    }
  }

  async function onStop() {
    setLoading(true);
    setError(null);
    try {
      await stopSession();
      await refresh();
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err));
    } finally {
      setLoading(false);
    }
  }

  async function togglePolicyLock() {
    // Removed policy lock control in simplified UI
  }

  return (
    <div className="page">
      <header className="topbar">
        <div>
          <h1>Activity Presence Controller</h1>
          <p className="muted">Local-first dashboard to manage presence sessions and policy locks.</p>
        </div>
        <span className={`badge ${stateLabel.toLowerCase()}`}>{stateLabel}</span>
      </header>

      {error && <div className="alert">{error}</div>}

      <div className="grid">
        <section className="card">
          <header>
            <h2>Controls</h2>
            <button className="ghost" onClick={refresh} disabled={loading}>Refresh</button>
          </header>

          <div className="actions">
            <button onClick={onStart} disabled={loading}>Start</button>
            <button onClick={onStop} className="ghost" disabled={loading}>Stop</button>
          </div>
          <div className="help">Start will wait ~2s then keep the cursor moving across the screen to stay active.</div>
        </section>

        <section className="card">
          <header>
            <h2>Status</h2>
          </header>
          <div className="status-grid">
            <div>
              <p className="label">Agent state</p>
              <p className="value">{status ? status.state : "..."}</p>
            </div>
            <div>
              <p className="label">Mode</p>
              <p className="value">{status?.mode || "-"}</p>
            </div>
            <div>
              <p className="label">Policy lock</p>
              <p className="value">{status?.policyLocked ? "Enabled" : "Disabled"}</p>
            </div>
            <div>
              <p className="label">Next action</p>
              <p className="value">{status?.nextAction || "-"}</p>
            </div>
          </div>
        </section>

        <section className="card logs">
          <header>
            <h2>Action Log</h2>
          </header>
          <div className="log-list">
            {logs.length === 0 && <p className="muted">No entries yet.</p>}
            {logs.map((entry: AgentLogEntry) => (
              <div key={entry.ts} className="log-item">
                <div className="log-ts">{new Date(entry.ts).toLocaleString()}</div>
                <div className="log-action">{entry.action}</div>
              </div>
            ))}
          </div>
        </section>
      </div>
    </div>
  );
}

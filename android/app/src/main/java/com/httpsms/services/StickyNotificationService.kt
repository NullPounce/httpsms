package com.httpsms.services

import android.app.Service
import android.content.Intent
import android.os.IBinder
import timber.log.Timber

@Suppress("DEPRECATION")
class StickyNotificationService : Service() {

    override fun onBind(intent: Intent?): IBinder? {
        Timber.d("Some component want to bind with the service [${intent?.action}]")
        return null
    }

    override fun onCreate() {
        Timber.d("The service has been created")
        super.onCreate()
    }

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        Timber.d("onStartCommand executed with startId: $startId")



        // If we return START_STICKY, the service will be restarted if it gets terminated by the system
        return START_STICKY
    }

    override fun onDestroy() {
        super.onDestroy()
        Timber.d("The service has been destroyed")
        // Remove the notification when the service is destroyed
        stopForeground(true)
    }

}

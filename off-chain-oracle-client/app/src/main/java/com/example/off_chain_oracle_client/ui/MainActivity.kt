package com.example.off_chain_oracle_client.ui

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.Button
import androidx.activity.viewModels
import com.example.off_chain_oracle_client.R
import com.example.off_chain_oracle_client.model.Validity
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class MainActivity : AppCompatActivity() {

    private val viewModel: MainViewModel by viewModels()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val btnValid = findViewById<Button>(R.id.valid_btn)
        val btnInvalid = findViewById<Button>(R.id.invalid_btn)


        btnValid.setOnClickListener {
            viewModel.setValidityStateEvent(ValidityStateEvent.ValidityEvent(Validity(true)))
        }

        btnInvalid.setOnClickListener {
            viewModel.setValidityStateEvent(ValidityStateEvent.ValidityEvent(Validity(false)))
        }


    }
}
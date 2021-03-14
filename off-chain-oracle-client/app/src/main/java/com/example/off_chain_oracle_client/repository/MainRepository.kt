package com.example.off_chain_oracle_client.repository

import com.example.off_chain_oracle_client.model.Validity
import com.example.off_chain_oracle_client.network.NetworkApi

class MainRepository (
    private val networkApi: NetworkApi
) {

    suspend fun setValidity(validity: Validity) {
        networkApi.postValidity(validity)
    }

}
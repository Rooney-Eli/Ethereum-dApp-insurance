package com.example.off_chain_oracle_client.network

import com.example.off_chain_oracle_client.model.Validity
import retrofit2.http.Body
import retrofit2.http.POST

interface NetworkApi {

    @POST("/")
    suspend fun postValidity(@Body valid: Validity)
}